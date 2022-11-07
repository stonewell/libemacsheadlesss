#ifndef __HEADLESS__H__
#define __HEADLESS__H__

#include "headless_service.h"

void bridge_keyboard_input(keyboard_input input, uint32_t keycode);
void bridge_new_client(new_client nc, uint32_t clientId);
void bridge_client_disconnect(client_disconnect cd, uint32_t clientId);

#endif //__HEADLESS__H__
