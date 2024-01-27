file = ARGS[1]
(@ccall "export.so".cmd_gat(file::Cstring)::Cstring)
