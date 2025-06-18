import pandas as pd
import matplotlib.pyplot as plt
import os

# Locate csv
csv_path = os.path.join("results", "benchmark.csv")

# Read CSV file
df = pd.read_csv(csv_path, names=["Method", "Size", "Time", "Length"])
df["Time"] = pd.to_timedelta(df["Time"])

# convert time to millisecond
df["Time(ms)"] = df["Time"].dt.total_seconds() * 1000

# Pivot to make chartable
df["Size"] = pd.Categorical(df["Size"], categories=["1KB", "1MB", "100MB"], ordered=True)
pivot = df.pivot(index="Size", columns="Method", values="Time(ms)")
pivot = pivot.reindex(["1KB", "1MB", "100MB"]) # ensure order consistency

# Plot the data
ax = pivot.plot(kind="bar", ylabel="Time (ms)", title="AES-256 Encryption Benchmark", figsize=(8,5), colormap="viridis")
ax.set_yscale("log")
plt.xticks(rotation=0)
plt.tight_layout()

# Save it
output_path = os.path.join("results", "charts.png")
plt.savefig(output_path)
plt.show()
plt.show()