#include "headless_service.h"
#include "libheadless.h"

void
remote_cursor_to (int terminal_id, int vpos, int hpos)
{
}
void
remote_raw_cursor_to (int terminal_id, int row, int col)
{
}

void
remote_clear_to_end (int terminal_id)
{
}

void
remote_clear_frame (int terminal_id)
{
}

void
remote_clear_end_of_line (int terminal_id, int first_unused_hpos)
{
}

void
remote_ins_del_lines (int terminal_id, int vpos, int n)
{
}

void
remote_insert_glyphs (int terminal_id, const wchar_t *start, int len)
{
}

void
remote_write_glyphs (int terminal_id, const wchar_t *string, int len)
{
}

void
remote_delete_glyphs (int terminal_id, int n)
{
}

void
remote_ring_bell (int terminal_id)
{
}

void
remote_reset_terminal_modes (int terminal_id)
{
}

void
remote_set_terminal_modes (int terminal_id)
{
}

void
remote_update_end (int terminal_id)
{
}

void
remote_menu_show (int terminal_id, int x, int y, int menuflags,
                  const char * title, const char **error_name)
{
}

void
remote_set_terminal_window (int terminal_id, int size)
{
}

int
remote_defined_color (int terminal_id, const char *color_name,
                      const char *color_def, int alloc, int _makeIndex)
{
  return 0;
}

int
remote_read_avail_input (int terminal_id)
{
  return 0;
}

void
remote_delete_frame (int terminal_id)
{
}

void
remote_delete_terminal (int terminal_id)
{
}
