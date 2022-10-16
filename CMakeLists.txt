cmake_minimum_required(VERSION 3.10)
project(libnvimheadless C CXX)

if(POLICY CMP0075)
  cmake_policy(SET CMP0075 NEW)
endif()

set_property(GLOBAL PROPERTY USE_FOLDERS ON)

# used for check_c_compiler_flag
include(CheckCCompilerFlag)
include(CheckCXXCompilerFlag)

list(APPEND CMAKE_MODULE_PATH "${PROJECT_SOURCE_DIR}/neovim/cmake")

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set(CMAKE_C_STANDARD 11)
set(CMAKE_C_STANDARD_REQUIRED ON)

#set(CMAKE_SHARED_LINKER_FLAGS "-Wl,--no-undefined")

set(CMAKE_POSITION_INDEPENDENT_CODE ON)

if(CMAKE_BUILD_TYPE MATCHES Debug)
  set(DEBUG 1)
else()
  set(DEBUG 0)
endif()

include(CheckCSourceCompiles)

execute_process(COMMAND make libnvim
  WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/neovim
  RESULT_VARIABLE neovim_setup)

if (NOT neovim_setup EQUAL "0")
  message(FATAL_ERROR "failed to configure neovim:${neovim_setup}")
endif()

add_subdirectory(headless)
add_subdirectory(test)
