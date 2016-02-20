package downloader_test

import (
	. "github.com/adria-stef/TvShowDownloader/downloader"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Downloader", func() {
	Describe("Extracter", func() {
		Context("ExtractSeasonEpisode with", func() {
			It("valid input", func() {
				season, episode := ExtractSeasonEpisode("s01e13")
				Expect(season).To(Equal(1))
				Expect(episode).To(Equal(13))
			})
		})

		It("invalid Input", func() {
			season, episode := ExtractSeasonEpisode("batman")
			Expect(season).To(BeZero())
			Expect(episode).To(BeZero())

		})
	})

	Context("ExtarctShow with ", func() {
		It("valid input", func() {
			name, se := ExtarctShow("Impractical Jokers S05E02 HDTV x264-W4F.torrent")
			Expect(name).To(Equal("impractical jokers "))
			Expect(se).To(Equal("s05e02"))
		})

		It("invalid input", func() {
			name, se := ExtarctShow("batman")
			Expect(name).To(BeEmpty())
			Expect(se).To(BeEmpty())

		})
	})
})
