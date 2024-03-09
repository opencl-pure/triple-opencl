package constants

const (
	CL_DEVICE_HALF_FP_CONFIG                      = 0x1033
	cl_APPLE_SetMemObjectDestructor               = 1
	cl_APPLE_ContextLoggingFunctions              = 1
	cl_khr_icd                                    = 1
	CL_PLATFORM_ICD_SUFFIX_KHR                    = 0x0920
	CL_PLATFORM_NOT_FOUND_KHR                     = -1001
	CL_CONTEXT_MEMORY_INITIALIZE_KHR              = 0x2030
	CL_DEVICE_TERMINATE_CAPABILITY_KHR            = 0x2031
	CL_CONTEXT_TERMINATE_KHR                      = 0x2032
	cl_khr_terminate_context                      = 1
	CL_DEVICE_SPIR_VERSIONS                       = 0x40E0
	CL_PROGRAM_BINARY_TYPE_INTERMEDIATE           = 0x40E1
	CL_DEVICE_COMPUTE_CAPABILITY_MAJOR_NV         = 0x4000
	CL_DEVICE_COMPUTE_CAPABILITY_MINOR_NV         = 0x4001
	CL_DEVICE_REGISTERS_PER_BLOCK_NV              = 0x4002
	CL_DEVICE_WARP_SIZE_NV                        = 0x4003
	CL_DEVICE_GPU_OVERLAP_NV                      = 0x4004
	CL_DEVICE_KERNEL_EXEC_TIMEOUT_NV              = 0x4005
	CL_DEVICE_INTEGRATED_MEMORY_NV                = 0x4006
	cl_amd_device_memory_flags                    = 1
	CL_MEM_USE_PERSISTENT_MEM_AMD                 = 1 << 6 //AllocfromGPU'sCPUvisibleheap
	CL_DEVICE_MAX_ATOMIC_COUNTERS_EXT             = 0x4032
	CL_DEVICE_PROFILING_TIMER_OFFSET_AMD          = 0x4036
	CL_DEVICE_TOPOLOGY_AMD                        = 0x4037
	CL_DEVICE_BOARD_NAME_AMD                      = 0x4038
	CL_DEVICE_GLOBAL_FREE_MEMORY_AMD              = 0x4039
	CL_DEVICE_SIMD_PER_COMPUTE_UNIT_AMD           = 0x4040
	CL_DEVICE_SIMD_WIDTH_AMD                      = 0x4041
	CL_DEVICE_SIMD_INSTRUCTION_WIDTH_AMD          = 0x4042
	CL_DEVICE_WAVEFRONT_WIDTH_AMD                 = 0x4043
	CL_DEVICE_GLOBAL_MEM_CHANNELS_AMD             = 0x4044
	CL_DEVICE_GLOBAL_MEM_CHANNEL_BANKS_AMD        = 0x4045
	CL_DEVICE_GLOBAL_MEM_CHANNEL_BANK_WIDTH_AMD   = 0x4046
	CL_DEVICE_LOCAL_MEM_SIZE_PER_COMPUTE_UNIT_AMD = 0x4047
	CL_DEVICE_LOCAL_MEM_BANKS_AMD                 = 0x4048
	CL_DEVICE_THREAD_TRACE_SUPPORTED_AMD          = 0x4049
	CL_DEVICE_GFXIP_MAJOR_AMD                     = 0x404A
	CL_DEVICE_GFXIP_MINOR_AMD                     = 0x404B
	CL_DEVICE_AVAILABLE_ASYNC_QUEUES_AMD          = 0x404C
	CL_DEVICE_TOPOLOGY_TYPE_PCIE_AMD              = 1
	CL_HSA_ENABLED_AMD                            = 1 << 62
	CL_HSA_DISABLED_AMD                           = 1 << 63
	CL_CONTEXT_OFFLINE_DEVICES_AMD                = 0x403F
	CL_PRINTF_CALLBACK_ARM                        = 0x40B0
	CL_PRINTF_BUFFERSIZE_ARM                      = 0x40B1
	cl_ext_device_fission                         = 1
	CL_DEVICE_PARTITION_EQUALLY_EXT               = 0x4050
	CL_DEVICE_PARTITION_BY_COUNTS_EXT             = 0x4051
	CL_DEVICE_PARTITION_BY_NAMES_EXT              = 0x4052
	CL_DEVICE_PARTITION_BY_AFFINITY_DOMAIN_EXT    = 0x4053
	CL_DEVICE_PARENT_DEVICE_EXT                   = 0x4054
	CL_DEVICE_PARTITION_TYPES_EXT                 = 0x4055
	CL_DEVICE_AFFINITY_DOMAINS_EXT                = 0x4056
	CL_DEVICE_REFERENCE_COUNT_EXT                 = 0x4057
	CL_DEVICE_PARTITION_STYLE_EXT                 = 0x4058
	CL_IMAGE_BYTE_PITCH_AMD                       = 0x4059
	CL_DEVICE_PARTITION_FAILED_EXT                = -1057
	CL_INVALID_PARTITION_COUNT_EXT                = -1058
	CL_INVALID_PARTITION_NAME_EXT                 = -1059
	CL_AFFINITY_DOMAIN_L1_CACHE_EXT               = 0x1
	CL_AFFINITY_DOMAIN_L2_CACHE_EXT               = 0x2
	CL_AFFINITY_DOMAIN_L3_CACHE_EXT               = 0x3
	CL_AFFINITY_DOMAIN_L4_CACHE_EXT               = 0x4
	CL_AFFINITY_DOMAIN_NUMA_EXT                   = 0x10
	CL_AFFINITY_DOMAIN_NEXT_FISSIONABLE_EXT       = 0x100
	CL_PROPERTIES_LIST_END_EXT                    = 0
	CL_PARTITION_BY_COUNTS_LIST_END_EXT           = 0
	CL_PARTITION_BY_NAMES_LIST_END_EXT            = -1
	CL_MEM_EXT_HOST_PTR_QCOM                      = 1 << 29
	CL_DEVICE_EXT_MEM_PADDING_IN_BYTES_QCOM       = 0x40A0
	CL_DEVICE_PAGE_SIZE_QCOM                      = 0x40A1
	CL_IMAGE_ROW_ALIGNMENT_QCOM                   = 0x40A2
	CL_IMAGE_SLICE_ALIGNMENT_QCOM                 = 0x40A3
	CL_MEM_HOST_UNCACHED_QCOM                     = 0x40A4
	CL_MEM_HOST_WRITEBACK_QCOM                    = 0x40A5
	CL_MEM_HOST_WRITETHROUGH_QCOM                 = 0x40A6
	CL_MEM_HOST_WRITE_COMBINING_QCOM              = 0x40A7
	CL_MEM_ION_HOST_PTR_QCOM                      = 0x40A8
	CL_MEM_BUS_ADDRESSABLE_AMD                    = 1 << 30
	CL_MEM_EXTERNAL_PHYSICAL_AMD                  = 1 << 31
	CL_COMMAND_WAIT_SIGNAL_AMD                    = 0x4080
	CL_COMMAND_WRITE_SIGNAL_AMD                   = 0x4081
	CL_COMMAND_MAKE_BUFFERS_RESIDENT_AMD          = 0x4082
	cl_khr_sub_groups                             = 1
	CL_KERNEL_MAX_SUB_GROUP_SIZE_FOR_NDRANGE_KHR  = 0x2033
	CL_KERNEL_SUB_GROUP_COUNT_FOR_NDRANGE_KHR     = 0x2034
)
