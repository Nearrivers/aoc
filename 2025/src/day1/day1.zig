const std = @import("std");
const t = std.testing;

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

        const lineNumber = getLineNumber(line);

        const result = part2Parsing(dialPointedNumber, passwordCount, lineNumber);
        passwordCount = result.passwordCount;
        dialPointedNumber = result.dialPointedNumber;
    }

    std.debug.print("\nPassword is {d}\n", .{passwordCount});
}

fn getLineNumber(line: []const u8) i16 {
    const lineNumber: i16 = std.fmt.parseInt(i16, line[1..], 10) catch {
        std.debug.print("error: line {s} doesn't contain a number\n", .{line});
        std.process.exit(1);
    };

    if (std.mem.indexOf(u8, line, "L") != null) {
        return lineNumber * -1;
    }

    return lineNumber;
}

fn part2Parsing(dialValue: i16, pwdCount: i16, lineNumber: i16) struct { dialPointedNumber: i16, passwordCount: i16 } {
    var dialPointedNumber = dialValue;
    var passwordCount = pwdCount;

    const minimum: i16 = 0;
    const maximum: i16 = 100;

    const isDialPointingAtZeroBeforeTurning = dialPointedNumber == 0;

    dialPointedNumber += lineNumber;

    if (minimum < dialPointedNumber and dialPointedNumber < maximum) {
        return .{ .dialPointedNumber = dialPointedNumber, .passwordCount = passwordCount };
    }

    if (dialPointedNumber == 100 or dialPointedNumber == 0 or dialPointedNumber == -100) {
        dialPointedNumber = 0;
        passwordCount += 1;
        return .{ .dialPointedNumber = dialPointedNumber, .passwordCount = passwordCount };
    }

    var times0WasHit = @divFloor(dialPointedNumber, maximum);

    if (times0WasHit < 0) {
        times0WasHit *= -1;
    }

    if (isDialPointingAtZeroBeforeTurning and times0WasHit > 0 and lineNumber < maximum) {
        times0WasHit -= 1;
    }

    passwordCount += times0WasHit;

    const minusMaximum: i16 = maximum * -1;

    if (dialPointedNumber < 0 and dialPointedNumber > minusMaximum) {
        dialPointedNumber = maximum + dialPointedNumber;

        if (@abs(dialPointedNumber) == 100 or dialPointedNumber == 0) {
            dialPointedNumber = 0;
            passwordCount += 1;
        }

        return .{ .dialPointedNumber = dialPointedNumber, .passwordCount = passwordCount };
    } else if (dialPointedNumber < 0 and dialPointedNumber < minusMaximum) {
        dialPointedNumber = maximum + @rem(dialPointedNumber, minusMaximum);

        if (@abs(dialPointedNumber) == 100 or dialPointedNumber == 0) {
            dialPointedNumber = 0;
            passwordCount += 1;
        }

        return .{ .dialPointedNumber = dialPointedNumber, .passwordCount = passwordCount };
    }

    const remainder = @rem(dialPointedNumber, maximum);

    dialPointedNumber = remainder;
    return .{ .dialPointedNumber = dialPointedNumber, .passwordCount = passwordCount };
}

test "get line number" {
    const cases = [_]struct {
        line: []const u8,
        expectedOutcome: i16,
    }{
        .{ .line = "R1000", .expectedOutcome = 1000 },
        .{ .line = "L1000", .expectedOutcome = -1000 },
        .{ .line = "L1", .expectedOutcome = -1 },
        .{ .line = "R1", .expectedOutcome = 1 },
        .{ .line = "R13", .expectedOutcome = 13 },
        .{ .line = "L67", .expectedOutcome = -67 },
    };

    for (cases) |case| {
        const result = getLineNumber(case.line);
        try t.expectEqual(case.expectedOutcome, result);
    }
}

test "part 2 parsing" {
    const cases = [_]struct { input: []const u8, expectedPwdCount: i16, expectedDial: i16 }{
        .{ .input = "R1000,", .expectedPwdCount = 10, .expectedDial = 50 },
        .{ .input = "R50,L150,", .expectedPwdCount = 2, .expectedDial = 50 },
        .{ .input = "R50,", .expectedPwdCount = 1, .expectedDial = 0 },
        .{ .input = "L50,", .expectedPwdCount = 1, .expectedDial = 0 },
        .{ .input = "R50,R100,R150,", .expectedPwdCount = 3, .expectedDial = 50 },
        .{ .input = "R50,L107,R8,L2,", .expectedPwdCount = 4, .expectedDial = 99 },
        .{ .input = "L250,", .expectedPwdCount = 3, .expectedDial = 0 },
        .{ .input = "L251,", .expectedPwdCount = 3, .expectedDial = 99 },
        .{ .input = "R75,L25,R10,", .expectedPwdCount = 2, .expectedDial = 10 },
        .{ .input = "R30,L130,", .expectedPwdCount = 1, .expectedDial = 50 },
        .{ .input = "R391,L283,R769,L301", .expectedPwdCount = 18, .expectedDial = 26 },
        .{ .input = "R78,R165,R900", .expectedPwdCount = 11, .expectedDial = 93 },
        .{ .input = "L51", .expectedPwdCount = 1, .expectedDial = 99 },
        .{ .input = "L137,L672,L147", .expectedPwdCount = 10, .expectedDial = 94 },
        .{ .input = "L49,L1,R1", .expectedPwdCount = 1, .expectedDial = 1 },
        .{ .input = "R49,R1,L1", .expectedPwdCount = 1, .expectedDial = 99 },
    };

    for (cases) |case| {
        var lines = std.mem.splitScalar(u8, case.input, ',');
        var dialPointedNumber: i16 = 50;
        var passwordCount: i16 = 0;

        while (lines.next()) |raw_line| {
            const line = trimNewline(raw_line);

            if (line.len == 0) {
                break;
            }

            const lineNumber = getLineNumber(line);

            const result = part2Parsing(dialPointedNumber, passwordCount, lineNumber);
            passwordCount = result.passwordCount;
            dialPointedNumber = result.dialPointedNumber;
        }

        try t.expectEqual(case.expectedPwdCount, passwordCount);
        try t.expectEqual(case.expectedDial, dialPointedNumber);
    }
}
