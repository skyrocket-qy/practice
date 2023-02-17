def BFS(G, start):
    visited, queue, out = set([start]), [start], []
    while queue:
        start = queue.pop()
        for w in sorted(G[start]):
            if w not in visited:
                out.append((start, w))
                visited.add(w)
                queue.append(w)
    return out