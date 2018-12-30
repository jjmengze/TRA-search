package search

import (
	"github.com/TRA-search/pkg/search/util"
	"github.com/spf13/cobra"
)

type SearchOptions struct {
	FromCity        string
	FromStation     string
	FromStationName string
	ToCity          string
	ToStation       string
	ToStationName   string
	TrainClass      string
	searchdate      string
	FromTimeSelect  string
	ToTimeSelect    string
	Timetype        string
}

func NewSearchOptions(streams *util.IOStreams) *SearchOptions {
	return &SearchOptions{}
}

func NewCmdSearch(streams *util.IOStreams) *cobra.Command {
	o := NewSearchOptions(streams)

	cmd := &cobra.Command{
		Use:                   "search --fcity=cityname ",
		DisableFlagsInUseLine: true,
		Short:                 "Search a particular date",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	addRunFlags(cmd, o)

	return cmd
}

func addRunFlags(cmd *cobra.Command, opt *SearchOptions) {
	cmd.Flags().StringVar(&opt.FromCity, "fcity", "", "Where city are you from")
	cmd.Flags().StringVar(&opt.FromStation, "fstation", "", "Where station are you from")
	cmd.Flags().StringVar(&opt.ToCity, "tcity", "", "Where station are you from")
	cmd.Flags().StringVar(&opt.ToStation, "tstation", "", "Where city are you from")
	cmd.Flags().StringVar(&opt.TrainClass, "class", "", "Where city are you from")
	cmd.Flags().StringVar(&opt.searchdate, "date", "", "Where station are you from")
	cmd.Flags().StringVar(&opt.FromTimeSelect, "ftime", "", "Where city are you from")
	cmd.Flags().StringVar(&opt.ToTimeSelect, "ttime", "", "Where station are you from")

	cmd.MarkFlagRequired("fcity")
}
