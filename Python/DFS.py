def DFS(G, y):
    result, visited, lst = [y], set([y]), []
    while len(result) > 0:
        current = result.pop()
        for i in sorted(G[current]):
            if i not in visited:
                result.append(i)
                visited.add(i)
                lst.append((current, i))
                break
    return lst