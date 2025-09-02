import autoprefixer from 'autoprefixer';
import nesting from 'postcss-nesting';

export default {
	syntax: 'postcss-scss',
	plugins: [nesting(), autoprefixer()]
};
