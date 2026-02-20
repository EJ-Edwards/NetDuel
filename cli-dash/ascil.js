function art1(){
    console.log("  _   _      _   _               _ ");
    console.log(" | \\ | | ___| |_| |__   ___   __| |");
    console.log(" |  \\| |/ _ \\ __| '_ \\ / _ \\ / _` |");
    console.log(" | |\\  |  __/ |_| | | | (_) | (_| |");
    console.log(" |_| \\_|\\___|\\__|_| |_|\\___/ \\__,_|");
    console.log("");
}

function rock(){
    console.log("    _______");
    console.log("---'   ____)");
    console.log("      (_____");
    console.log("      (_____");
    console.log("      (____)");
    console.log("---.__(___)");
}

function paper(){
    console.log("    _______");
    console.log("---'   ____)____");
    console.log("          ______)");
    console.log("          _______)");
    console.log("         _______)");
    console.log("---.__________)");
}

function scissors(){
    console.log("    _______");
    console.log("---'   ____)____");
    console.log("          ______)");
    console.log("       __________)");
    console.log("      (____)");
    console.log("---.__(___)");
}

module.exports = { art1, rock, paper, scissors };

