const TelegramBot = require('node-telegram-bot-api');

const token = '5765052522:AAEdf9UfDLwrAPJUTfax7LAPHSo81kmY_fo';

const bot = new TelegramBot(token, {polling: true});

bot.on('message', async (msg) => {
  const chatId = msg.chat.id;
  const text = msg.text;

  if (text === '/start') {
    bot.sendMessage(chatId, "1", {
      reply_markup: {
        inline_keyboard: [
            [{text: 'Кнопка',
              web_app: {url: "https://yandex.ru"}}]
        ]
      }
    })
  }
});