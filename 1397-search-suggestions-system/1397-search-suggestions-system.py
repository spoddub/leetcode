class Node:
    def __init__(self):
        self.children = {}
        self.words = []

class Trie:

    def __init__(self):
        self.root = Node()

    def insert(self, word: str) -> None:
        node = self.root
        
        for ch in word:
            if ch not in node.children:
                node.children[ch] = Node()
            
            node = node.children[ch]
            node.words.append(word)

    def search(self, prefix: str) -> List[str]:
        node = self.root
        
        for ch in prefix:
            if ch not in node.children:
                return []
            node = node.children[ch]
        
        return node.words


class Solution:
    def suggestedProducts(self, products: List[str], searchWord: str) -> List[List[str]]:
        root = Trie()
        
        for word in products:
            root.insert(word)

        result: List[List[str]] = []
        
        for i in range(len(searchWord)):
            found = root.search(searchWord[:i + 1])
            result.append(sorted(found)[:3])
        
        return result