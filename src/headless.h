#ifndef __HEADLESS__H__
#define __HEADLESS__H__

#include <stdint.h>

typedef void (*keyboard_input)(uint32_t keycode);
typedef void (*new_client)(uint32_t clientId);

typedef struct tagClientCallbacks {
  keyboard_input input;
} ClientCallbacks;

typedef struct tagServerConfig {
  uint32_t port;
  const char * addr;

  new_client new_client_callback;
} ServerConfig;

void bridge_keyboard_input(keyboard_input input, uint32_t keycode);
void bridge_new_client(new_client nc, uint32_t clientId);

#endif //__HEADLESS__H__
