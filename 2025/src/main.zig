const std = @import("std");
const cli = @import("cli/cli.zig");
const day1 = @import("day1/day1.zig");

// Usage : zig run main.zig -- <day> <part> <flow>
pub fn main() !void {
    const allocator = std.heap.page_allocator;
    var aocCli = try cli.parseArgs(allocator);

    aocCli.printArgs();

    const fullPath = std.fmt.allocPrint(allocator, "./day{d}/{s}.txt", .{ aocCli.day, @tagName(aocCli.flow) }) catch |err| {
        std.debug.print("error: could not determine file path: {}", .{err});
        std.process.exit(1);
    };

    defer allocator.free(fullPath);

    var challengeFile = std.fs.cwd().openFile(fullPath, .{ .mode = .read_only }) catch |err| {
        std.debug.print("error: unable to open challenge file: {}", .{err});
        std.process.exit(1);
    };

    defer challengeFile.close();

    const st = try challengeFile.stat();
    if (st.kind != .file) {
        std.debug.print("error: challenge element is not a file", .{});
        std.process.exit(1);
    }

    switch (aocCli.day) {
        1 => day1.day1(challengeFile, allocator),
        else => {
            std.debug.print("error: day not done yet\n", .{});
            std.process.exit(1);
        },
    }
}
