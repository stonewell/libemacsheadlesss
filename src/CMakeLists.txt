set(HEADLESS_SERV headless-serv)

file(GLOB HEADLESS_SERV_SOURCES RELATIVE "${CMAKE_CURRENT_SOURCE_DIR}" "*.go")

set (RPC_SOURCES ${CMAKE_CURRENT_LIST_DIR}/../proto/emacsheadless.pb.go
  ${CMAKE_CURRENT_LIST_DIR}/../proto/emacsheadless_grpc.pb.go
)

add_custom_command(OUTPUT ${CMAKE_CURRENT_BINARY_DIR}/${HEADLESS_SERV}
  COMMAND env GOPATH=${GOPATH} ${CMAKE_Go_COMPILER} build
  -buildmode=exe
  -o ${CMAKE_CURRENT_BINARY_DIR}/${HEADLESS_SERV}
  ${CMAKE_CURRENT_SOURCE_DIR}
  DEPENDS   gen_rpc
    ${HEADLESS_SERV_SOURCES}
    ${RPC_SOURCES}
  WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
  VERBATIM
)

add_custom_target(${HEADLESS_SERV} ALL DEPENDS
  ${CMAKE_CURRENT_BINARY_DIR}/${HEADLESS_SERV}
)

add_dependencies(${HEADLESS_SERV} gen_rpc)

install(PROGRAMS ${CMAKE_CURRENT_BINARY_DIR}/${HEADLESS_SERV} DESTINATION bin)

add_custom_command(OUTPUT ${CMAKE_CURRENT_BINARY_DIR}/libheadless.so
  COMMAND env GOPATH=${GOPATH} ${CMAKE_Go_COMPILER} build
  -buildmode=c-shared
  -o ${CMAKE_CURRENT_BINARY_DIR}/libheadless.so
  ${CMAKE_CURRENT_SOURCE_DIR}
  DEPENDS   gen_rpc
    ${HEADLESS_SERV_SOURCES}
    ${RPC_SOURCES}
  WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
  VERBATIM
)

add_custom_target(libheadless ALL DEPENDS
  ${CMAKE_CURRENT_BINARY_DIR}/libheadless.so
)

add_dependencies(libheadless gen_rpc)

install(PROGRAMS ${CMAKE_CURRENT_BINARY_DIR}/libheadless.so DESTINATION lib)
