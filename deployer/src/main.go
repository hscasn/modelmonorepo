package main

import (
	"deployer/cloudrun"
	"deployer/utils"
	"deployer/versionsets"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

var acceptedEnvironments = []string{
	"prod",
}

type args struct {
	Environment string
	Version     string
}

type deployer interface {
	Deploy() error
	Describe() string
}

type deployPlan struct {
	CloudRun []cloudrun.Config
}

func main() {
	args := utils.Must(parseArgs()).(args)

	// Getting versionset
	vs := parseVersionset(args.Version)

	// Building deployment plans
	plans := deployPlan{
		CloudRun: cloudrun.MakeDeployPlans(args.Environment, vs.CloudRun),
	}

	// Deploying
	deployAll(plans)
}

func deployAll(plans deployPlan) error {
	v := reflect.ValueOf(plans)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		deployables := v.Field(i)
		platform := t.Field(i).Name
		for k := 0; k < deployables.Len(); k++ {
			deployable := deployables.Index(k).Interface().(deployer)
			fmt.Printf("Deploying %s %s", platform, deployable.Describe())
			if err := deployable.Deploy(); err != nil {
				return fmt.Errorf("could not deploy %s %s", platform, deployable.Describe())
			}
		}
	}
	return nil
}

func parseVersionset(version string) versionsets.Versionset {
	vsbt := utils.Must(readVersionsetFile(version)).([]byte)
	return utils.Must(versionsets.Parse(vsbt)).(versionsets.Versionset)
}

func readVersionsetFile(version string) ([]byte, error) {
	return ioutil.ReadFile(fmt.Sprintf("../versionsets/%s.yaml", version))
}

func parseArgs() (args, error) {
	environment := os.Args[1]
	version := os.Args[2]
	return args{
		Environment: environment,
		Version:     version,
	}, nil
}
