import EditorJS from '@editorjs/editorjs';
import Header from '@editorjs/header';
import List from '@editorjs/list';
import Delimiter from '@editorjs/delimiter';
import Paragraph from '@editorjs/paragraph';
import Embed from '@editorjs/embed';
import SimpleImage from '@editorjs/simple-image';

$(function() {
    let data = {};
    if($('#editContainer').length){
        data = $('#content').val()
        data = JSON.parse(data)
    }

    const editor = new EditorJS({
            holder: 'editorjs',
            /**
             * Available Tools list.
             * Pass Tool's class or Settings object for each Tool you want to use
             */
            tools:{
                header:Header,
                delimiter: Delimiter,
                paragraph: {
                    class: Paragraph,
                    inlineToolbar: true,
                },
                embed: Embed,
                image: SimpleImage,
            },
            data:data
        }
    );

    $( "#saveButton" ).click(function() {
        editor.save().then((output) => {
            $("#content").val(JSON.stringify(output))
        }).catch((error) => {
            console.log('Saving failed: ', error)
        });
    });
});