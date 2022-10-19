#include "headless.h"

__attribute__ ((visibility ("hidden")))
void bridge_keyboard_input(keyboard_input input, uint32_t keycode)
{
  if (input) {
    input(keycode);
  }
}

__attribute__ ((visibility ("hidden")))
void bridge_new_client(new_client nc, uint32_t clientId)
{
  if (nc)
    nc(clientId);
}
