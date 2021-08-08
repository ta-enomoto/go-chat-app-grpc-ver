const path = require('path');

module.exports = {
  mode: 'development',
  entry: {
    adminchatroom: './web-server/src/adminchatroom.js',
    chatroom: './web-server/src/chatroom.js'
  },
  output: {
    filename: '[name].js',
    path: path.join(__dirname, 'web-server/public/scripts')
  }
};