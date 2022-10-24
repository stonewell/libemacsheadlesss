set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_CURRENT_SOURCE_DIR}/cmake")

cmake_minimum_required(VERSION 3.0)

project(emacsheadless Go C CXX)

include(cmake/golang.cmake)
include(cmake/flags.cmake)

add_subdirectory(proto)
add_subdirectory(pkg)
add_subdirectory(internal)
add_subdirectory(test)
