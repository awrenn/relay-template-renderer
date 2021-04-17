package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/puppetlabs/relay-sdk-go/pkg/outputs"
	"github.com/puppetlabs/relay-sdk-go/pkg/taskutil"
)

type TemplateSpec struct {
	Template   string `spec:"template"`
	Parameters string `spec:"parameters"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	defaultMetadataSpecURL, err := taskutil.MetadataSpecURL()
	if err != nil {
		return err
	}

	// This seems like it could done better
	specURL := flag.String("spec-url", defaultMetadataSpecURL, "url to fetch spec from")
	flag.Parse()

	planOpts := taskutil.DefaultPlanOptions{SpecURL: *specURL}
	spec := TemplateSpec{}
	if err := taskutil.PopulateSpecFromDefaultPlan(&spec, planOpts); err != nil {
		return err
	}
	log.Printf("DEBUG: %+v", planOpts)

	oc, err := outputs.NewDefaultOutputsClientFromNebulaEnv()
	if err != nil {
		return err
	}

	ctx, cls := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cls()
	if err := oc.SetOutput(ctx, "output", "Hello world!"); err != nil {
		return err
	}
	return nil
}
