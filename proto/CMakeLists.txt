execute_process(COMMAND
  go env GOPATH
  OUTPUT_VARIABLE FOO
  OUTPUT_STRIP_TRAILING_WHITESPACE)

add_custom_command(OUTPUT ${CMAKE_CURRENT_LIST_DIR}/emacsheadless.pb.go
    ${CMAKE_CURRENT_LIST_DIR}/emacsheadless_grpc.pb.go
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
  ${CMAKE_CURRENT_LIST_DIR}/emacsheadless.pb.go
  ${CMAKE_CURRENT_LIST_DIR}/emacsheadless_grpc.pb.go
)
