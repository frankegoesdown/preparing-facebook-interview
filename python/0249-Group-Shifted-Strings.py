import collections


def groupStrings(strings):
    groups = collections.defaultdict(list)
    print(strings)
    for s in strings:
        print(ord(s[0]))
        print([ord(c) - ord(s[0]) % 26 for c in s])
        print(groups)
        print("="*8)
        groups[tuple((ord(c) - ord(s[0])) % 26 for c in s)].append(s),
    return groups.values()


print(groupStrings(["abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"]))
