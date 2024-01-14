package template

import "github.com/wildegor/daily_tasks/go/internal/data_structures/hash_set"

/**
There are n rooms labeled from 0 to n - 1 and all the rooms are locked except for room 0. Your goal is to visit all the rooms. However, you cannot enter a locked room without having its key.
When you visit a room, you may find a set of distinct keys in it. Each key has a number on it, denoting which room it unlocks, and you can take all of them with you to unlock the other rooms.
Given an array rooms where rooms[i] is the set of keys that you can obtain if you visited room i, return true if you can visit all the rooms, or false otherwise.
@link https://leetcode.com/problems/keys-and-rooms/
*/

func canVisitAllRooms(rooms [][]int) bool {
	visited := hash_set.NewHashSet()

	var dfs func(int)
	dfs = func(room int) {
		visited.Add(room)
		for _, key := range rooms[room] {
			if !visited.Has(key) {
				dfs(key)
			}
		}
	}

	dfs(0)

	return visited.Size() == len(rooms)
}

var Solution = canVisitAllRooms
