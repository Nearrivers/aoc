const std = @import("std");

pub const Flow = enum { example, input };

pub const AocCli = struct {
    flow: Flow = .example,
    day: u8 = 0,
    part: u8 = 0,

    pub fn printArgs(self: *AocCli) void {
        std.debug.print("flow {s} - day {} - part {}\n", .{ @tagName(self.flow), self.day, self.part });
    }
};

pub fn printUsage() void {
    std.debug.print("usage: zig run main.zig -- <day> <part> <flow>\n", .{});
}

pub fn parseArgs(allocator: std.mem.Allocator) !AocCli {
    var cli: AocCli = .{};
    const args = try std.process.argsAlloc(allocator);
    defer std.process.argsFree(allocator, args);

    if (args.len != 4) {
        printUsage();
        std.process.exit(0);
    }

    const maybeDay = try allocator.dupe(u8, args[1]);
    cli.day = std.fmt.parseInt(u8, maybeDay, 10) catch {
        std.debug.print("error: day flag should be a number\n", .{});
        std.process.exit(1);
    };

    const maybePart = try allocator.dupe(u8, args[2]);
    cli.part = std.fmt.parseInt(u8, maybePart, 10) catch {
        std.debug.print("error: part flag should be a number\n", .{});
        std.process.exit(1);
    };

    const maybeFlow = try allocator.dupe(u8, args[3]);

    if (std.mem.indexOf(u8, maybeFlow, @tagName(.example)) != null) {
        cli.flow = .example;
    } else if (std.mem.indexOf(u8, maybeFlow, @tagName(.input)) != null) {
        cli.flow = .input;
    } else {
        printUsage();
        std.debug.print("error: flow flag unknown : got {s}, expected 'input' or 'example'\n", .{maybeFlow});
        std.process.exit(1);
    }

    return cli;
}
