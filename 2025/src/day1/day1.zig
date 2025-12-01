const std = @import("std");
const cli = @import("../cli/cli.zig");

fn trimNewline(line: []const u8) []const u8 {
    if (line.len > 0 and line[line.len - 1] == '\r') {
        return line[0 .. line.len - 1];
    }
    return line;
}

pub fn day1(challengeFile: std.fs.File, allocator: std.mem.Allocator, part: u8) !void {
    const max_bytes = 8 * 1024 * 1024;
    const contents = challengeFile.readToEndAlloc(allocator, max_bytes) catch |err| switch (err) {
        error.FileTooBig => {
            std.debug.print("error: file exceeds {} bytes limit\n", .{max_bytes});
            std.process.exit(1);
        },
        else => return err,
    };
    defer allocator.free(contents);

    var lines = std.mem.splitScalar(u8, contents, '\n');
    switch (part) {
        1 => part1(&lines),
        2 => part2(&lines),
        else => {
            std.debug.print("error: unexpected 'part' flag value. The part flag can only be 1 or 2", .{});
            std.process.exit(1);
        },
    }
}

fn part1(lines: *std.mem.SplitIterator(u8, .scalar)) void {
    var passwordCount: i16 = 0;
    var dialPointedNumber: i16 = 50;
    while (lines.next()) |raw_line| {
        const line = trimNewline(raw_line);

        if (line.len == 0) {
            break;
        }

        std.debug.print("\nline {s}\n", .{line});

        const lineNumber = std.fmt.parseInt(i16, line[1..], 10) catch {
            std.debug.print("error: line {s} doesn't contain a number\n", .{line});
            std.process.exit(1);
        };

        std.debug.print("line number: {d}\n", .{lineNumber});

        const minimum: i16 = 0;
        const maximum: i16 = 100;

        const modulo = @mod(lineNumber, maximum);

        if (std.mem.indexOf(u8, line, "L") != null) {
            dialPointedNumber -= modulo;
        } else {
            dialPointedNumber += modulo;
        }

        std.debug.print("dialPointerNumber is {d}\n", .{dialPointedNumber});

        if (dialPointedNumber == 100 or dialPointedNumber == 0) {
            dialPointedNumber = 0;
            passwordCount += 1;
            continue;
        }

        if (minimum < dialPointedNumber and dialPointedNumber < maximum) {
            continue;
        }

        if (dialPointedNumber < 0) {
            dialPointedNumber = maximum + dialPointedNumber;
            std.debug.print("dialPointerNumber has been changed to {d}\n", .{dialPointedNumber});
            continue;
        }

        const remainder = @rem(dialPointedNumber, maximum);
        dialPointedNumber = remainder;
    }

    std.debug.print("\nPassword is {d}\n", .{passwordCount});
}

fn part2(lines: *std.mem.SplitIterator(u8, .scalar)) void {
    var passwordCount: i16 = 0;
    var dialPointedNumber: i16 = 50;

    while (lines.next()) |raw_line| {
        const line = trimNewline(raw_line);

        if (line.len == 0) {
            break;
        }

        std.debug.print("\nline {s}\n", .{line});

        const lineNumber = std.fmt.parseInt(i16, line[1..], 10) catch {
            std.debug.print("error: line {s} doesn't contain a number\n", .{line});
            std.process.exit(1);
        };

        std.debug.print("line number: {d}\n", .{lineNumber});

        const minimum: i16 = 0;
        const maximum: i16 = 100;

        const modulo = @mod(lineNumber, maximum);

        if (std.mem.indexOf(u8, line, "L") != null) {
            dialPointedNumber -= modulo;
        } else {
            dialPointedNumber += modulo;
        }

        std.debug.print("dialPointerNumber is {d}\n", .{dialPointedNumber});

        if (dialPointedNumber == 100 or dialPointedNumber == 0) {
            dialPointedNumber = 0;
            passwordCount += 1;
            continue;
        }

        if (minimum < dialPointedNumber and dialPointedNumber < maximum) {
            continue;
        }

        if (dialPointedNumber < 0) {
            dialPointedNumber = maximum + dialPointedNumber;
            std.debug.print("dialPointerNumber has been changed to {d}\n", .{dialPointedNumber});
            continue;
        }

        const remainder = @rem(dialPointedNumber, maximum);
        dialPointedNumber = remainder;
    }

    std.debug.print("\nPassword is {d}\n", .{passwordCount});
}
