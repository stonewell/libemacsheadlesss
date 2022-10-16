find_package(LuaJit REQUIRED)
find_package(LibUV REQUIRED)
find_package(Msgpack REQUIRED)

set(nvim_headless_src
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/api/ui.c

  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/api/private/helpers.c

  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/ui_bridge.c
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/highlight.c
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/memory.c
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/math.c
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/map.c
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/event/multiqueue.c
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/decoration_provider.c
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/main.c
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/ex_cmds.c
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/ui.c
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/fileio.c
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/quickfix.c
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/mark.c

  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/os/env.c

  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/msgpack_rpc/server.c
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src/nvim/msgpack_rpc/channel.c
)

add_library(nvimheadless SHARED
  ${nvim_headless_src})

target_include_directories(nvimheadless PRIVATE
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/build/cmake.config
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/build/include
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/build/src/nvim/auto
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src
  ${LUAJIT_INCLUDE_DIRS}
  ${MSGPACK_INCLUDE_DIRS}
)

target_compile_definitions(nvimheadless PUBLIC
  EXITFREE
  INCLUDE_GENERATED_DECLARATIONS
  MVIM_MSGPACK_HAS_FLOAT32
  NVIM_UNIBI_HAS_VAR_FROM
)

target_link_libraries(nvimheadless
  ${LUAJIT_LIBRARY}
  ${MSGPACK_LIBRARY}
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/.deps/build/src/libuv-build/libuv_a.a
  -lpthread
  -ldl
  -lrt
  -lnsl
  -lm
  -lutil
  -fstack-protector-strong
)
