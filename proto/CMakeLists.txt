execute_process(COMMAND
  go env GOPATH
  OUTPUT_VARIABLE FOO
  OUTPUT_STRIP_TRAILING_WHITESPACE)

set(RPC_SOURCES ${CMAKE_CURRENT_LIST_DIR}/emacsheadless.pb.go
  ${CMAKE_CURRENT_LIST_DIR}/emacsheadless_grpc.pb.go
)

set(RPC_SOURCES ${RPC_SOURCES} PARENT_SCOPE)

add_custom_command(OUTPUT ${RPC_SOURCES}
  COMMAND env GOPATH=${GOPATH} PATH=$ENV{PATH}:${FOO}/bin protoc
    --go_out=${CMAKE_CURRENT_LIST_DIR}
    --go_opt=paths=source_relative
    --go-grpc_out=${CMAKE_CURRENT_LIST_DIR}
    --go-grpc_opt=paths=source_relative
    -I ${CMAKE_CURRENT_LIST_DIR}
    ${CMAKE_CURRENT_LIST_DIR}/emacsheadless.proto
  WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
  DEPENDS ${CMAKE_CURRENT_LIST_DIR}/emacsheadless.proto
  VERBATIM
)

add_custom_target(gen_rpc ALL DEPENDS
  ${RPC_SOURCES}
)
