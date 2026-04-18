import * as dotenv from 'dotenv';
import * as path from 'path';

const pathEnv = process.env.NODE_ENV === 'production' ? '.env.production' : '.env.development';

export default dotenv.config({
     path: path.resolve(process.cwd(), pathEnv),
})
