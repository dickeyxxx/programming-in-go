// Copyright © 2010-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"math"
	"sort"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
	stdDev  float64
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = Mean(numbers)
	stats.median = Median(numbers)
	stats.stdDev = StdDev(numbers)
	return stats
}

func Mean(numbers []float64) float64 {
	return Sum(numbers) / float64(len(numbers))
}

func Sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func Median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

func StdDev(numbers []float64) float64 {
	var nums []float64
	var n float64
	avg := Mean(numbers)
	for _, x := range numbers {
		nums = append(nums, math.Pow(x-avg, 2))
	}
	n = Sum(nums)
	n = n / float64(len(nums)-1)
	n = math.Sqrt(n)
	return n
}
