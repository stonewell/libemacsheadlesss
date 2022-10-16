add_executable(test_headless
  test.c
  test_main.cpp)

set_target_properties(test_headless PROPERTIES LINKER_LANGUAGE CXX)

target_link_libraries(test_headless PRIVATE
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/build/lib/libnvim.a
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/.deps/usr/lib/libunibilium.a
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/.deps/usr/lib/libtermkey.a
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/.deps/usr/lib/libluv.a
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/.deps/usr/lib/libuv_a.a
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/.deps/usr/lib/libluajit-5.1.a
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/.deps/usr/lib/libvterm.a
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/.deps/usr/lib/libmsgpackc.a
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/.deps/usr/lib/libtree-sitter.a
  -lpthread
  -ldl
  -lrt
  -lnsl
  -lm
  -lutil)

target_include_directories(test_headless PRIVATE
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/build/cmake.config
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/build/include
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/build/src/nvim/auto
  ${CMAKE_CURRENT_SOURCE_DIR}/../neovim/src
)
