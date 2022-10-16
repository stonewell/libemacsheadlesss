#pragma once

#ifdef __cplusplus
extern "C"
{
#endif

  typedef struct nvim_main_s * nvim_main_t;

  nvim_main_t run_nvim_main();

  void stop_nvim_main(nvim_main_t n);

#ifdef __cplusplus
}
#endif
