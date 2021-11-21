// Imports
const { Client, Intents } = require('discord.js');

const sleep = function(t) {
    return new Promise((res, _) => {
        setTimeout(res, t);
    });
}; 

const main = async function () {
    const client = new Client({ intents: [ Intents.FLAGS.GUILDS ] });

    // register events
    client.prependOnceListener('ready', () => {
        console.log('Logged in as %s', client.user.username);        
        
        let guild = client.guilds.cache.get('718993047563534398');
        let role = guild.roles.cache.find(r => r.name.toLowerCase() == 'members');

        // for async (let member of (await guild.members.fetch())) {
        //     console.log(member.user);
        //     if (member.user.bot) continue;
        //     if (member.roles.cache.has(role.id)) continue;
            
        //     member.roles.add(role);

        //     console.log('Added %s to %s', role.name, member.displayName);
        //     // sleep(2.5 * 1_000);
        // }
    });

    client.prependListener('message', () => {
        return;
    });

    client.login('ODA2MTc5MDU5NTA2NzQxMjY4.YBlqwQ.rTs_BtjMmgQytfGPx0yVvmcdJPU');
};

main()
    .catch(e => console.error(`Main error: ${e}`));