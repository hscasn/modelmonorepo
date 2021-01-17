package cloudrun

import (
	"deployer/utils"
	"deployer/versionsets"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Config is the main struct for Cloud Run configuration
type Config struct {
	ModuleName           string `yaml:"module_name"`
	ModuleVersion        string
	Project              string            `yaml:"project"`
	Region               string            `yaml:"region"`
	AllowUnauthenticated bool              `yaml:"allow_unauthenticated"`
	SQLInstances         []string          `yaml:"sql_instances"`
	MaxInstances         int               `yaml:"max_instances"`
	ServiceAccount       string            `yaml:"service_account"`
	Env                  map[string]string `yaml:"env"`
}

type prod struct {
	Prod Config `yaml:"prod"`
}

func parseYaml(environment string, yamlConfig []byte, moduleVersion string) Config {
	t := prod{}

	err := yaml.Unmarshal(yamlConfig, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	t.Prod.ModuleVersion = moduleVersion

	return t.Prod
}

// Describe cloudrun module
func (cfg Config) Describe() string {
	return fmt.Sprintf("%s:%s", cfg.ModuleName, cfg.ModuleVersion)
}

// Deploy cloudrun module
func (cfg Config) Deploy() error {
	moduleName := strings.Replace(cfg.ModuleName, "_", "-", -1)
	arg := []string{
		fmt.Sprintf("--project=%s", cfg.Project),
		"beta",
		"run",
		"deploy",
		moduleName,
		"--region", cfg.Region,
		"--quiet",
		"--image", fmt.Sprintf("gcr.io/%s/%s:%s", cfg.Project, cfg.ModuleName, cfg.ModuleVersion),
		"--service-account", cfg.ServiceAccount,
		"--platform", "managed",
		"--max-instances", fmt.Sprintf("%d", cfg.MaxInstances),
	}

	if cfg.AllowUnauthenticated {
		arg = append(arg, "--allow-unauthenticated")
	}

	numSQLInstances := 0
	sqlInstances := ""
	if len(cfg.SQLInstances) > 0 {
		arg = append(arg, "--add-cloudsql-instances")
		for _, instance := range cfg.SQLInstances {
			if numSQLInstances > 0 {
				sqlInstances += ","
			}
			sqlInstances += instance
			numSQLInstances++
		}
		arg = append(arg, sqlInstances)
	}

	if len(cfg.Env) > 0 {
		arg = append(arg, "--set-env-vars")
		envVars := ""
		numEnvVars := 0
		for key, val := range cfg.Env {
			if numEnvVars > 0 {
				envVars += ","
			}
			envVars += fmt.Sprintf("%s=%s", key, val)
			numEnvVars++
		}
		arg = append(arg, envVars)
	}

	cmd := exec.Command("gcloud", arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error deploying module %s: %w", cfg.Describe(), err)
	}
	return nil
}

// MakeDeployPlans will parse all configurations to deploy the module and build a struct with all this info
func MakeDeployPlans(environment string, vs versionsets.PlatformVersionset) []Config {
	r := make([]Config, 0)
	for moduleName, moduleVersion := range vs {
		bt := utils.Must(utils.ReadModuleFile(moduleName)).([]byte)
		cfg := parseYaml(environment, bt, moduleVersion)
		r = append(r, cfg)
	}
	return r
}
