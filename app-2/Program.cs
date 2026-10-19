using System.Text;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;

var factory = new ConnectionFactory
{
    Uri = new Uri("amqp://admin:admin@rabbitmq:5672/")
};

using var connection = await factory.CreateConnectionAsync();
using var channel = await connection.CreateChannelAsync();

Console.WriteLine("conectado ao rabbitmq com sucesso.");

await channel.QueueDeclareAsync(
    queue: "fila-pizzas",
    durable: true,
    exclusive: false,
    autoDelete: false,
    arguments: null
);

var consumer = new AsyncEventingBasicConsumer(channel);

consumer.ReceivedAsync += (model, ea) =>
{
    var body = ea.Body.ToArray();
    var mensagem = Encoding.UTF8.GetString(body);
    Console.WriteLine($"pizza recebida: {mensagem}");
    return Task.CompletedTask;
};

await channel.BasicConsumeAsync(
    queue: "fila-pizzas",
    autoAck: true,
    consumer: consumer
);

await Task.Delay(Timeout.Infinite);
