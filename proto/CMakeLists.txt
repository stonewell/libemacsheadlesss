execute_process(COMMAND
  go env GOPATH
  OUTPUT_VARIABLE FOO
  OUTPUT_STRIP_TRAILING_WHITESPACE)

message("${FOO}")

add_custom_command(OUTPUT ${OUTPUT_DIR}/.timestamp
  COMMAND env GOPATH=${GOPATH} PATH=$ENV{PATH}:${FOO}/bin protoc --go_out=${CMAKE_CURRENT_BINARY_DIR}
  --go_opt=paths=source_relative
  --go-grpc_out=${CMAKE_CURRENT_BINARY_DIR}
  --go-grpc_opt=paths=source_relative
  -I ${CMAKE_CURRENT_LIST_DIR}
  ${CMAKE_CURRENT_LIST_DIR}/emacsheadless.proto
  WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
)

add_custom_target(gen_rpc ALL DEPENDS ${OUTPUT_DIR}/.timestamp
  ${CMAKE_CURRENT_LIST_DIR}/emacsheadless.proto
)
