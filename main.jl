using IOCapture: IOCapture
using TerminalPager: pager

function gat(file::AbstractString)
    (@ccall "export.so".cmd_gat(file::Cstring)::Cstring)
    nothing
end

function gess(file::AbstractString)
    c = IOCapture.capture() do
        gat(file)
    end

    out = c.output
    out |> pager
end

if abspath(PROGRAM_FILE) == @__FILE__
    if length(ARGS) != 1
        println("Usage: julia $(PROGRAM_FILE) <filename>")
        exit(1)
    end
    file = ARGS[begin]
    gat(file)
end
