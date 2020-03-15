	defaultFilePathPrefix                  = "/var/lib/m3db"
	defaultWriteBufferSize                 = 65536
	defaultDataReadBufferSize              = 65536
	defaultInfoReadBufferSize              = 128
	defaultSeekReadBufferSize              = 4096
	defaultThroughputLimitMbps             = 100.0
	defaultThroughputCheckEvery            = 128
	defaultForceIndexSummariesMmapMemory   = false
	defaultForceBloomFilterMmapMemory      = false
	defaultBloomFilterFalsePositivePercent = 0.02

	// BloomFilterFalsePositivePercent controls the target false positive percentage
	// for the bloom filters for the fileset files.
	BloomFilterFalsePositivePercent *float64 `yaml:"bloomFilterFalsePositivePercent"`
	if f.BloomFilterFalsePositivePercent != nil &&
		(*f.BloomFilterFalsePositivePercent < 0 || *f.BloomFilterFalsePositivePercent > 1) {
		return fmt.Errorf(
			"fs bloomFilterFalsePositivePercent is set to: %f, but must be between 0.0 and 1.0",
			*f.BloomFilterFalsePositivePercent)
	}
// BloomFilterFalsePositivePercentOrDefault returns the configured value for the target
// false positive percent for the bloom filter for the fileset files if configured, or a default
// value otherwise
func (f FilesystemConfiguration) BloomFilterFalsePositivePercentOrDefault() float64 {
	if f.BloomFilterFalsePositivePercent != nil {
		return *f.BloomFilterFalsePositivePercent
	}
	return defaultBloomFilterFalsePositivePercent
}
