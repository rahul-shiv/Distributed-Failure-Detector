import numpy as np
import matplotlib.pyplot as plt
infrequent = np.array([138.27,144.93,138.27,131.33,131.91])
frequent = np.array([161.77,174.31,111.53,111.58,110.99])
regex = np.array([205,195,192,190,195])

# Calculate the average
infrequent_mean = np.mean(infrequent)
frequent_mean = np.mean(frequent)
regex_mean = np.mean(regex)

# Calculate the standard deviation
infrequent_std = np.std(infrequent)
frequent_std = np.std(frequent)
regex_std = np.std(regex)

# Create lists for the plot
labels = ['Infrequent("cupiditate")', 'Frequent ("distributed")', 'Regex ("dist.*")']
x_pos = np.arange(len(labels))
CTEs = [infrequent_mean, frequent_mean, regex_mean]
error = [infrequent_std, infrequent_std, regex_std]

# Build the plot
fig, ax = plt.subplots()
ax.bar(x_pos, CTEs, yerr=error, align='center', alpha=0.5, ecolor='black', capsize=10)
ax.set_ylabel('Time (ms)')
ax.set_xticks(x_pos)
ax.set_xticklabels(labels)
ax.set_title('End to end Latency of distributed grep')
ax.yaxis.grid(True)

# Save the figure and show
plt.tight_layout()
plt.savefig('bar_plot_with_error_bars.png')
plt.show()