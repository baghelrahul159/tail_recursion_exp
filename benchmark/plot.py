import matplotlib.pyplot as plt

ns = [1, 2, 5, 10, 20, 50, 100, 300, 500, 700, 1000]
data = {
  'recursionWithMemoization': [3.00, 8.03, 10.8, 25.9, 20.2, 27.7, 33.4, 28.7, 19.7, 21.3, 16.3],
  'tail_recursion': [4.33, 7.42, 24.2, 69.8, 126, 270, 477, 1368, 3005, 4191, 5076],
  'iterative': [2.52, 2.40, 6.86, 12.0, 20.1, 59.4, 68.1, 194, 359, 398, 1038],
}
for (label, values) in data.items():
    plt.plot(ns, values, label=label)
plt.legend()
plt.yscale('log')
plt.show()