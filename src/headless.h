#ifndef __HEADLESS__H__
#define __HEADLESS__H__

#include <stdint.h>

typedef void (*keyboard_input)(uint32_t keycode);
typedef struct tagServerCallbacks {
  keyboard_input input;
} ServerCallbacks;

void bridge_keyboard_input(keyboard_input input, uint32_t keycode);

#endif //__HEADLESS__H__
