import extractorSvelte from '@unocss/extractor-svelte';
import presetIcons from '@unocss/preset-icons';
import transformerDirectives from '@unocss/transformer-directives';
import { defineConfig, presetUno } from 'unocss';

export default defineConfig({
	presets: [
		presetUno(),
		presetIcons({
			extraProperties: {
				display: 'inline-block',
				'vertical-align': 'middle'
			}
		})
	],
	extractors: [extractorSvelte()],
	transformers: [
		transformerDirectives({
			applyVariable: ['--apply']
		})
	]
});
