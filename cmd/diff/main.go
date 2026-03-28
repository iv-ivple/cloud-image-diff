package main
 
import (
    "context"
    "encoding/json"
    "fmt"
    "os"
 
    "github.com/spf13/cobra"
    "github.com/iv-ivple/cloud-image-diff/internal/image"
    "github.com/iv-ivple/cloud-image-diff/internal/providers"
)
 
var (release, arch, outputFmt string; awsRegions []string)
 
func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err); os.Exit(1)
    }
}
 
var rootCmd = &cobra.Command{
    Use:   "cloud-image-diff",
    Short: "Compare Ubuntu cloud images across AWS, Azure, and GCP",
}
 
var listCmd = &cobra.Command{
    Use: "list", RunE: func(cmd *cobra.Command, args []string) error {
        p := providers.NewAWSProvider(awsRegions)
        imgs, err := p.ListImages(context.Background(),
            image.ImageFilter{Release: release, Arch: arch})
        if err != nil { return err }
        return printJSON(imgs)
    },
}
 
var diffCmd = &cobra.Command{
    Use: "diff <providerA> <providerB>", Args: cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
        filter := image.ImageFilter{Release: release, Arch: arch}
        pA := providers.NewAWSProvider(awsRegions)
        pB := providers.NewAWSProvider(awsRegions)
        imgsA, _ := pA.ListImages(context.Background(), filter)
        imgsB, _ := pB.ListImages(context.Background(), filter)
        return printJSON(image.Diff(imgsA, imgsB))
    },
}
func printJSON(v any) error {
    enc := json.NewEncoder(os.Stdout)
    enc.SetIndent("", "  ")
    return enc.Encode(v)
}
 
func init() {
    rootCmd.PersistentFlags().StringVar(&release,"release","24.04","Ubuntu release")
    rootCmd.PersistentFlags().StringVar(&arch,"arch","amd64","Architecture")
    rootCmd.PersistentFlags().StringVar(&outputFmt,"output","json","json or table")
    listCmd.Flags().StringSliceVar(&awsRegions,"aws-regions",
        []string{"us-east-1","eu-west-1"},"AWS regions")
    diffCmd.Flags().StringSliceVar(&awsRegions,"aws-regions",
        []string{"us-east-1","eu-west-1"},"AWS regions")
    rootCmd.AddCommand(listCmd, diffCmd)
}
