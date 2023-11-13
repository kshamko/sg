package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/jessevdk/go-flags"
	"github.com/kshamko/sg/internal/database"
	"github.com/kshamko/sg/internal/model"
	"github.com/kshamko/sg/internal/predict"
)

const fileBasePath = "./assets/data/"

func main() {

	var opts struct {
		Source    string `long:"source" description:"source file with data" required:"true"`
		Aggregate string `long:"aggregate" description:"aggregation type" required:"true" choice:"country" choice:"campaign"`
		Model     string `long:"model" description:"model used for prediction" required:"true" choice:"linex" choice:"expsmoothing"`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		return
	}

	datasource, err := initDatasource(opts.Source)
	if err != nil {
		fmt.Printf("Error while parsing file %s: %s\n", opts.Source, err)
		os.Exit(1)
	}

	model, err := initModel(opts.Model)
	if err != nil {
		fmt.Printf("Error while loading model %s: %s\n", opts.Model, err)
		os.Exit(1)
	}

	predictor, err := predict.New(
		datasource,
		model,
		predict.NewDefaultAggregator(),
		opts.Aggregate,
	)

	if err != nil {
		fmt.Printf("Error while initialising predictor: %s\n", err)
		os.Exit(1)
	}
	preds := predictor.Predict(60)

	sort.Slice(preds, func(i, j int) bool { return preds[i].Key < preds[j].Key })

	for _, pred := range preds {
		fmt.Printf("%s: %f\n", pred.Key, pred.Value)
	}
}

func initDatasource(sourceFile string) (predict.Database, error) {

	filePath := fileBasePath + sourceFile

	ext := filepath.Ext(filePath)

	if ext == ".csv" {
		return database.NewCSV(filePath)
	}

	if ext == ".json" {
		return database.NewJSON(filePath)
	}

	return nil, fmt.Errorf("unknown file extension %s", ext)
}

func initModel(modelName string) (predict.Inferencer, error) {

	if modelName == "linex" {
		return model.NewLinExtrapol(), nil
	}

	if modelName == "expsmoothing" {
		return model.NewExpSmooth(0.1, 0.25, 0.35), nil
	}

	return nil, fmt.Errorf("unknown model %s", modelName)
}
