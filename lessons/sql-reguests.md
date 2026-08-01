Получить статистику, сгруппированную по месяцу

```sql
SELECT to_char(date, 'YYYY-MM') as period, sum(clicks) FROM stats
WHERE date BETWEEN '01/01/2026' and '01/01/2027'
GROUP BY period
ORDER BY period
```

Получить статистику, сгруппированную по дню

```sql
SELECT to_char(date, 'YYYY-MM-DD') as period, sum(clicks) FROM stats
WHERE date BETWEEN '01/01/2026' and '01/01/2027'
GROUP BY period
ORDER BY period
```
