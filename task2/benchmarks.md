## pkg: task2/pkg/sorted/list

|name                                | exec_count | speed          | B/op        |  memory_allocated  |
| ---------------------------------- | ---------- | -------------- | ----------- | ------------------ |
| BenchmarkInsertIntoTheBeginning-12 | 7430068    |	141 ns/op      | 48 B/op	 | 1 allocs/op        |
| BenchmarkInsertIntoMiddle-12       | 81074	  | 69711 ns/op    | 63 B/op	 | 2 allocs/op        |
| BenchmarkInsertIntoTheEnd-12       | 7643908	  | 158 ns/op      | 56 B/op	 | 2 allocs/op        |
| BenchmarkInsertBatch-12            | 31	      | 36830848 ns/op | 558013 B/op | 19744 allocs/op    |
| BenchmarkGetMinOn100-12            | 1000000000 | 1.14 ns/op     | 0 B/op	     | 0 allocs/op        |
| BenchmarkGetMinOn10000-12          | 999477498  | 1.26 ns/op     | 0 B/op	     | 0 allocs/op        |
| BenchmarkGetMaxOn100-12            | 980176653  | 1.12 ns/op     | 0 B/op	     | 0 allocs/op        |
| BenchmarkGetMaxOn10000-12          | 1000000000 | 1.09 ns/op     | 0 B/op	     | 0 allocs/op        |
| BenchmarkEqual-12                  | 169940	  | 6088 ns/op     | 0 B/op	     | 0 allocs/op        |
| BenchmarkNotEqualByLen-12          | 268851075  | 4.43 ns/op     | 0 B/op	     | 0 allocs/op        |
| BenchmarkNotEqualInTheEnd-12       | 147525576  | 8.09 ns/op     | 0 B/op	     | 0 allocs/op        |
| BenchmarkNotEqualInTheBegin-12     | 147384612  | 7.99 ns/op     | 0 B/op	     | 0 allocs/op        |
| BenchmarkNotEqualInMiddle-12       | 388100	  | 3322 ns/op     | 0 B/op	     | 0 allocs/op        |


## pkg: task2/pkg/sorted/slice

|name                                | exec_count | speed          | B/op        |  memory_allocated  |
| ---------------------------------- | ---------- | -------------- | ----------- | ------------------ |
| BenchmarkInsertIntoTheBeginning-12 | 35798	  | 31669 ns/op	   | 106500 B/op | 1 allocs/op        |
| BenchmarkInsertIntoMiddle-12       | 36780	  | 33205 ns/op	   | 106500 B/op | 1 allocs/op        |
| BenchmarkInsertIntoTheEnd-12       | 35764	  | 33404 ns/op	   | 106500 B/op | 1 allocs/op        |
| BenchmarkInsertBatch-12            | 40	      | 28791548 ns/op | 354306 B/op | 4 allocs/op        |
| BenchmarkDeleteFromBeginning-12    | 369196	  | 3215 ns/op	   | 0 B/op	     | 0 allocs/op        |
| BenchmarkDeleteFromMiddle-12       | 281426	  | 3591 ns/op	   | 0 B/op	     | 0 allocs/op        |
| BenchmarkDeleteFromEnd-12          | 251934	  | 4142 ns/op	   | 0 B/op	     | 0 allocs/op        |
| BenchmarkGetMinOn100-12            | 995474844  |	1.16 ns/op	   | 0 B/op	     | 0 allocs/op        |
| BenchmarkGetMinOn10000-12          | 1000000000 |	1.12 ns/op	   | 0 B/op	     | 0 allocs/op        |
| BenchmarkGetMaxOn100-12            | 1000000000 |	1.10 ns/op	   | 0 B/op	     | 0 allocs/op        |
| BenchmarkGetMaxOn10000-12          | 989529262  |	1.08 ns/op	   | 0 B/op	     | 0 allocs/op        |
| BenchmarkEqual-12                  | 302220	  | 4138 ns/op	   | 8192 B/op	 | 1 allocs/op        |
| BenchmarkNotEqualByLen-12          | 341287678  |	3.54 ns/op	   | 0 B/op	     | 0 allocs/op        |
| BenchmarkNotEqualInTheEnd-12       | 230861536  |	4.80 ns/op	   | 0 B/op	     | 0 allocs/op        |
| BenchmarkNotEqualInTheBegin-12     | 251399209  |	4.68 ns/op	   | 0 B/op	     | 0 allocs/op        |
| BenchmarkNotEqualInMiddle-12       | 327960	  | 3552 ns/op	   | 8192 B/op	 | 1 allocs/op        |