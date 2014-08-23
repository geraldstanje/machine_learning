package main

import (
  "testing"
  "errors"
  "fmt"
  "github.com/sridif/gosvm"
)

func TestSvmClassifier(t *testing.T) {
  model, err := gosvm.LoadModel("svm_model")
  if err != nil {
    t.Fatal(err)
  }

  label := model.Predict(gosvm.FromDenseVector([]float64{-1, 1}))
  if label != 1.0 {
    str := fmt.Sprintf("Wrong predicted label: %f\n", label)
    t.Fatal(errors.New(str))
  }

  label = model.Predict(gosvm.FromDenseVector([]float64{1, 1}))
  if label != -1.0 {
    str := fmt.Sprintf("Wrong predicted label: %f\n", label)
    t.Fatal(errors.New(str))
  }

  label = model.Predict(gosvm.FromDenseVector([]float64{1, -1}))
  if label != 1.0 {
    str := fmt.Sprintf("Wrong predicted label: %f\n", label)
    t.Fatal(errors.New(str))
  }

  label = model.Predict(gosvm.FromDenseVector([]float64{-1, -1}))
  if label != -1.0 {
    str := fmt.Sprintf("Wrong predicted label: %f\n", label)
    t.Fatal(errors.New(str))
  }
}