const { spawn } = require('child_process');
const readline = require('readline');
const path = require('path');
const { art1, rock, paper, scissors } = require('./ascil');
const { connectToServer } = require('./server');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

function runPythonGame() {
  const pythonScript = path.join(__dirname, '..', 'engine', 'main.py');
  const py = spawn('python', [pythonScript]);

  py.stdout.on('data', (data) => {
    process.stdout.write(data.toString());
  });

  py.stderr.on('data', (data) => {
    process.stderr.write(data.toString());
  });

  rl.on('line', (input) => {
    py.stdin.write(input + '\n');
  });

  py.on('close', (code) => {
    console.log(`Python game exited with code ${code}`);
    rl.close();
  });
}

function runMultiplayer() {
  console.log('Multiplayer Room Options:');
  console.log('1. Create Room');
  console.log('2. Join Room');
  rl.question('Choose an option (1 or 2): ', (opt) => {
    if (opt === '1') {
      const roomPin = Math.floor(1000 + Math.random() * 9000).toString();
      console.log('Your room pin is:', roomPin);
      startMultiplayerWithPin(roomPin);
    } else if (opt === '2') {
      rl.question('Enter room pin: ', (roomPin) => {
        startMultiplayerWithPin(roomPin);
      });
    } else {
      console.log('Invalid option. Returning to menu.');
      mainMenu();
    }
  });
}

function startMultiplayerWithPin(roomPin) {
  const ws = connectToServer(
    () => {
      rl.question('Enter your username: ', (username) => {
        ws.send(JSON.stringify({ type: 'join', username, info: roomPin }));
        promptMove(ws);
      });
    },
    (data) => {
      const msg = JSON.parse(data);
      if (msg.type === 'info') {
        console.log('[INFO]', msg.info);
      } else if (msg.type === 'result') {
        console.log('[RESULT]', msg.result);
      }
      promptMove(ws);
    },
    () => {
      console.log('Disconnected from server.');
      rl.close();
    }
  );
}

function promptMove(ws) {
  rl.question('Enter your move (rock, paper, scissors): ', (move) => {
    ws.send(JSON.stringify({ type: 'move', move }));
  });
}

function mainMenu() {
  art1();
  console.log('Welcome to NetDuel CLI!');
  console.log('1. Single Player (Python Engine)');
  console.log('2. Multiplayer (Go Server)');
  console.log('3. View Rules');
  console.log('4. Exit');
  rl.question('Enter your choice: ', (choice) => {
    switch (choice) {
      case '1':
        runPythonGame();
        break;
      case '2':
        runMultiplayer();
        break;
      case '3':
        console.log('\nRules:');
        console.log('1. Choose single player or multiplayer mode.');
        console.log('2. Each player starts with 3 lives.');
        console.log('3. On each turn, choose rock, paper, or scissors.');
        console.log('4. Winner of the round keeps their lives, loser loses one life.');
        console.log('5. If both choose the same, it\'s a tie and no lives are lost.');
        console.log('6. First to reduce the opponent to 0 lives wins!');
        mainMenu();
        break;
      case '4':
        console.log('Goodbye!');
        rl.close();
        process.exit(0);
        break;
      default:
        console.log('Invalid choice, please try again.');
        mainMenu();
    }
  });
}

mainMenu();
