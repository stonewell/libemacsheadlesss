#include <assert.h>
#include <inttypes.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#include "test_main.h"

#include "nvim/grid.h"
#include "nvim/drawscreen.h"

bool server_init(const char *listen_addr);
void event_init(void);
void normal_enter(bool cmdwin, bool noexmode);
void win_init_size(void);
int win_alloc_first(void);
void set_init_1(bool clean_arg);
void set_init_2(bool clean_arg);

int nvim_main(int argc, char **argv);  // silence -Wmissing-prototypes

int main(){
  nvim_main_t n;

  n = run_nvim_main();

  printf("%d:%d\n", default_grid.rows, default_grid.cols);

  grid_puts(&default_grid,
            "this is a test",
            1,
            1,
            0);

  stop_nvim_main(n);
  return 0;
}
