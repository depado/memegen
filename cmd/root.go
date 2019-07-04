package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AddMemeFlags adds support to configure the level of the logger
func AddMemeFlags(c *cobra.Command) {
	c.PersistentFlags().StringP("top", "t", "", "text to display on top")
	c.PersistentFlags().StringP("bottom", "b", "", "text to display on bottom")
	c.PersistentFlags().StringP("output", "o", "out.png", "output file")
	c.PersistentFlags().Float64P("font.size", "s", 70, "font size")
	c.PersistentFlags().Int("font.outline", 4, "outline of the font")
	c.PersistentFlags().StringP("font.path", "f", "fonts/impact.ttf", "font path to use")
}

// AddAllFlags will add all the flags provided in this package to the provided
// command and will bind those flags with viper
func AddAllFlags(c *cobra.Command) {
	AddMemeFlags(c)
	if err := viper.BindPFlags(c.PersistentFlags()); err != nil {
		logrus.WithError(err).WithField("step", "AddAllFlags").Fatal("Couldn't bind flags")
	}
}
