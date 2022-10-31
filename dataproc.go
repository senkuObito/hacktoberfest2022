import (
        "context"
        "fmt"
        "io"

        dataproc "cloud.google.com/go/dataproc/apiv1"
        "google.golang.org/api/option"
        dataprocpb "google.golang.org/genproto/googleapis/cloud/dataproc/v1"
)

func createCluster(w io.Writer, projectID, region, clusterName string) error {
        // projectID := "your-project-id"
        // region := "us-central1"
        // clusterName := "your-cluster"
        ctx := context.Background()

        // Create the cluster client.
        endpoint := region + "-dataproc.googleapis.com:443"
        clusterClient, err := dataproc.NewClusterControllerClient(ctx, option.WithEndpoint(endpoint))
        if err != nil {
                return fmt.Errorf("dataproc.NewClusterControllerClient: %v", err)
        }
        defer clusterClient.Close()

        // Create the cluster config.
        req := &dataprocpb.CreateClusterRequest{
                ProjectId: projectID,
                Region:    region,
                Cluster: &dataprocpb.Cluster{
                        ProjectId:   projectID,
                        ClusterName: clusterName,
                        Config: &dataprocpb.ClusterConfig{
                                MasterConfig: &dataprocpb.InstanceGroupConfig{
                                        NumInstances:   1,
                                        MachineTypeUri: "n1-standard-2",
                                },
                                WorkerConfig: &dataprocpb.InstanceGroupConfig{
                                        NumInstances:   2,
                                        MachineTypeUri: "n1-standard-2",
                                },
                        },
                },
        }

        // Create the cluster.
        op, err := clusterClient.CreateCluster(ctx, req)
        if err != nil {
                return fmt.Errorf("CreateCluster: %v", err)
        }

        resp, err := op.Wait(ctx)
        if err != nil {
                return fmt.Errorf("CreateCluster.Wait: %v", err)
        }

        // Output a success message.
        fmt.Fprintf(w, "Cluster created successfully: %s", resp.ClusterName)
        return nil
}
