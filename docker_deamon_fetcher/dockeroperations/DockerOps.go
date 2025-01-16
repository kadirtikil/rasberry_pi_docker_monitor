package dockeroperations


import (
    "fmt"
    
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
    "golang.org/x/net/context"
)

/**
    TODO: 
        - send refresh after executing one of these ops DONE (use Inspect)
*/


func RestartContainer(id string) (types.ContainerJSON, error) {
    cli, err := client.NewClientWithOpts(client.WithVersion("1.41"))
    if err != nil {
        return types.ContainerJSON{}, fmt.Errorf("Error trying to create the client: %v", err)
    }


    if err := cli.ContainerRestart(context.Background(), id, container.StopOptions{}); err != nil {
        return types.ContainerJSON{}, fmt.Errorf("Error trying to start the container: %v", err)
    }
    
    jsonContainer, err := inspectContainer(id)
    if err != nil {
        return types.ContainerJSON{}, fmt.Errorf("Error trying to inspect container after restarting it!") 
    }


    return jsonContainer, nil
}


func PauseContainer(id string) (types.ContainerJSON, error){
    cli, err := client.NewClientWithOpts(client.WithVersion("1.41"))
    if err != nil {
        return types.ContainerJSON{}, fmt.Errorf("Error trying to create the client in PauseContainer: %v", err)
    }

    if err := cli.ContainerPause(context.Background(), id); err != nil {
        return types.ContainerJSON{}, fmt.Errorf("Error trying to start the container: %v", err)
    }

    jsonContainer, err := inspectContainer(id)
    if err != nil {
        return types.ContainerJSON{}, fmt.Errorf("Error trying to inspect the container after pausing it: %v", err)
    }

    return jsonContainer, nil 
}


func RemoveContainer(id string) (types.ContainerJSON, error) {
    cli, err := client.NewClientWithOpts(client.WithVersion("1.41"))
    if err != nil {
        return types.ContainerJSON{}, fmt.Errorf("Error trying to create the client in PauseContainer: %v", err)
    }

    if err := cli.ContainerRemove(context.Background(), id, container.RemoveOptions{}); err != nil {
        return types.ContainerJSON{}, fmt.Errorf("Error trying to start the container: %v", err)
    }
    
    return types.ContainerJSON{} ,nil
}


func KillContainer(id string) (types.ContainerJSON, error) {
    cli, err := client.NewClientWithOpts(client.WithVersion("1.41"))
    if err!= nil {
        return types.ContainerJSON{}, fmt.Errorf("Error trying to create the client in KillContainer: %v", err)
    }

    // its 15 now but i should  change it later into taking the signal as an argument
    // such that it can be controlled as well for containers that for exmaple need to be shut down asap or something
    if err := cli.ContainerKill(context.Background(), id, "SIGKILL"); err != nil {
        return types.ContainerJSON{}, fmt.Errorf("Error trying to kill the container in KillContainer: %v", err)
    }
    

    jsonContainer, err := inspectContainer(id)
    if err != nil {
        return types.ContainerJSON{}, fmt.Errorf("Error trying to inspect the container after killing it: %v", err)
    }

    return jsonContainer, nil

}


func inspectContainer(id string) (types.ContainerJSON, error) {
    cli, err := client.NewClientWithOpts(client.WithVersion("1.41"))
    if err != nil {
        return types.ContainerJSON{}, fmt.Errorf("Error trying to create the client in InspectContainer: %v", err)
    }
   
    container, err := cli.ContainerInspect(context.Background(), id)
    if err != nil {
        return types.ContainerJSON{}, fmt.Errorf("Error trying to inspect the container in InspectContainer: %v", err)
    }

    return container, nil
}