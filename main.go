package main

import (
  "context"
  "log"
  "os"
  "strings"

  "github.com/docker/docker/api/types"
  "github.com/docker/docker/client"
)

const (
  SERVICE_ENV_VAR = "PLUGIN_SERVICE_NAME"
)


func main() {
  service_name := os.Getenv(SERVICE_ENV_VAR)
  if service_name == "" {
    log.Fatalf("%s env var is not defined", SERVICE_ENV_VAR)
    os.Exit(-1)
  }

  cli, err := client.NewEnvClient()
  if err != nil {
    log.Fatal(err)
    os.Exit(-1)
  }

  // Fetching service information
  service, _, err := cli.ServiceInspectWithRaw(context.Background(), service_name)
  if err != nil {
    log.Fatal(err)
    os.Exit(-1)
  }

  // Removing hash part of image name so Docker can update it
  image := service.Spec.TaskTemplate.ContainerSpec.Image
  if !strings.Contains(image, "@") {
    log.Fatalf("Service image name \"%s\" does not follow pattern " +
      "${IMAGE_NAME}@${HASH_TYPE}:${HASH}", image)
    os.Exit(-1)
  }
  service.Spec.TaskTemplate.ContainerSpec.Image = strings.Split(image, "@")[0]

  // Updating service
  _, err = cli.ServiceUpdate(context.Background(), service.ID, service.Meta.Version, service.Spec,
    types.ServiceUpdateOptions{})
  if err != nil {
    log.Fatal(err)
    os.Exit(-1)
  }

  log.Println("Done!")
}
