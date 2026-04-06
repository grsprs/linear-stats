# linear-stats

Reads numeric data from a file and computes the **Linear Regression Line** and **Pearson Correlation Coefficient**.

## Usage

```bash
go run main.go data.txt
```

## Testing

1. Download and give executable permissions to the provided `stat-bin` folder:
   ```bash
   chmod +x bin/linear-stats
   ```
2. Generate test data:
   ```bash
   ./bin/linear-stats
   ```
3. Run the program with the generated file:
   ```bash
   go run main.go data.txt
   ```
4. Compare the outputs — values should match.

## Output Format

```
Linear Regression Line: y = <6 decimal places>x + <6 decimal places>
Pearson Correlation Coefficient: <10 decimal places>
```
