#include <assert.h>
#include <inttypes.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <thread>

#include "test_main.h"

extern "C"
{
  int nvim_main(int argc, char **argv);  // silence -Wmissing-prototypes
}

struct nvim_main_s {
  std::thread t;
};

static
int nvim_main_thread_func() {
  /* event_init(); */

  /* env_init(); */
  /* estack_init(); */

  /* printf("win_alloc_first:%d\n", win_alloc_first()); */

  /* printf("%d\n", server_init("localhost:1234")); */

  /* win_init_size(); */

  /* set_init_1(false); */
  /* set_init_2(true); */

  /* normal_enter(false, false); */

  char * args[] = {
    "test",
    "-n",
    "-u",
    "tmp.lua",
    "--noplugin",
    "--headless",
    "--listen",
    "localhost:1234"
  };

  return nvim_main(sizeof(args) / sizeof(char *), args);
}

nvim_main_t run_nvim_main() {
  nvim_main_t n = new nvim_main_s;

  std::thread t(nvim_main_thread_func);

  n->t.swap(t);

  std::this_thread::sleep_for(std::chrono::milliseconds(5 * 1000));

  return n;
}

void stop_nvim_main(nvim_main_t n)
{
  n->t.join();
  delete n;
}
