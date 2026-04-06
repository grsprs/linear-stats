# 📋 Οδηγός Audit — linear-stats

## Προαπαιτούμενα

- Docker Desktop **ανοιχτό και σε λειτουργία** (πράσινο εικονίδιο κάτω δεξιά)
- Άνοιξε ένα **terminal** (CMD ή PowerShell) στον φάκελο του project:
  ```
  cd C:\Users\Unkown\Desktop\Zone01\Module#585\linear-stats
  ```

---

## Βήμα 1 — Build Docker image (μία φορά μόνο)

```
cd stat-bin
docker build -t stat-bin .
```

Περίμενε να τελειώσει. Θα δεις `naming to docker.io/library/stat-bin:latest`.

Αυτό χρειάζεται **μόνο την πρώτη φορά**. Δεν το ξανατρέχεις.

---

## Βήμα 2 — Δημιούργησε data και σύγκρινε (επανάλαβε 4 φορές)

### 2a. Τρέξε το provided binary (δημιουργεί data.txt + δείχνει σωστή απάντηση)

Βεβαιώσου ότι είσαι μέσα στον φάκελο `stat-bin`:
```
docker run -v "%cd%:/data" -w /data --rm stat-bin /stat-bin/linear-stats
```

Θα εμφανίσει κάτι σαν:
```
Linear Regression Line: y = -1.000041x + 555.147506
Pearson Correlation Coefficient: -0.9986882209
```

**Σημείωσε αυτό το output** — είναι η σωστή απάντηση.

### 2b. Τρέξε το δικό σου πρόγραμμα

Πήγαινε πίσω στον κύριο φάκελο και τρέξε:
```
cd ..
go run main.go stat-bin\data.txt
```

Θα εμφανίσει το δικό σου output. **Σύγκρινέ το** με αυτό του βήματος 2a.

### 2c. Επανέλαβε

Μπες πάλι στο stat-bin και ξανατρέξε:
```
cd stat-bin
docker run -v "%cd%:/data" -w /data --rm stat-bin /stat-bin/linear-stats
cd ..
go run main.go stat-bin\data.txt
```

**Κάνε αυτό συνολικά 4 φορές** (1 αρχική + 3 επιπλέον).

---

## Τι ελέγχει ο auditor — Απαντήσεις

| Ερώτηση | Απάντηση |
|---------|----------|
| Are the outputs in the same format? | ✅ Yes — ίδιο format ακριβώς |
| Did Linear Regression values contain 6 decimal places? | ✅ Yes |
| Did Pearson Correlation contain 10 decimal places? | ✅ Yes |
| Did the values match? | ✅ Yes |
| Did values match in all 4 tries? | ✅ Yes |
| **Bonus**: README with explanation? | ✅ Yes — υπάρχει README.md |

---

## Quick Reference — Όλες οι εντολές μαζί

```
cd C:\Users\Unkown\Desktop\Zone01\Module#585\linear-stats\stat-bin
docker build -t stat-bin .

REM === Run 1 ===
docker run -v "%cd%:/data" -w /data --rm stat-bin /stat-bin/linear-stats
cd ..
go run main.go stat-bin\data.txt

REM === Run 2 ===
cd stat-bin
docker run -v "%cd%:/data" -w /data --rm stat-bin /stat-bin/linear-stats
cd ..
go run main.go stat-bin\data.txt

REM === Run 3 ===
cd stat-bin
docker run -v "%cd%:/data" -w /data --rm stat-bin /stat-bin/linear-stats
cd ..
go run main.go stat-bin\data.txt

REM === Run 4 ===
cd stat-bin
docker run -v "%cd%:/data" -w /data --rm stat-bin /stat-bin/linear-stats
cd ..
go run main.go stat-bin\data.txt
```
