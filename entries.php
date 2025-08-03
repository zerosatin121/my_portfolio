<?php
require 'config.php';

// Fetch logs from DB
$stmt = $pdo->query("SELECT * FROM logs ORDER BY log_date DESC");
$logs = $stmt->fetchAll(PDO::FETCH_ASSOC);
?>

<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <title>My Daily Logs</title>
  <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-gray-50 font-mono p-6">
  <h1 class="text-3xl font-bold mb-6">📚 My Daily Logs</h1>

  <?php if (count($logs) === 0): ?>
    <p class="text-gray-600">No entries yet. Start logging your progress!</p>
  <?php else: ?>
    <div class="space-y-6">
      <?php foreach ($logs as $log): ?>
        <div class="border p-4 rounded bg-white shadow">
          <div class="flex justify-between items-center mb-2">
            <span class="text-sm text-gray-500"><?= htmlspecialchars($log['log_date']) ?></span>
            <span class="text-xs px-2 py-1 rounded bg-purple-100 text-purple-700"><?= htmlspecialchars($log['category']) ?></span>
          </div>
          <p class="mb-2 whitespace-pre-line"><?= nl2br(htmlspecialchars($log['thoughts'])) ?></p>
          <?php if ($log['tools_used']): ?>
            <p class="text-sm text-gray-600">🛠 Tools: <?= htmlspecialchars($log['tools_used']) ?></p>
          <?php endif; ?>
          <?php if ($log['challenge_day']): ?>
            <p class="text-sm text-gray-600">📅 Day: <?= htmlspecialchars($log['challenge_day']) ?></p>
          <?php endif; ?>
        </div>
      <?php endforeach; ?>
    </div>
  <?php endif; ?>

  <a href="log.html" class="mt-10 inline-block text-sm text-gray-600 underline hover:text-black">← Back to Log Entry</a>
</body>
</html>
