import EditorJS from '@editorjs/editorjs';
import Header from '@editorjs/header';
import List from '@editorjs/list';
import Delimiter from '@editorjs/delimiter';
import Paragraph from '@editorjs/paragraph';
import Embed from '@editorjs/embed';
import CodeTool from '/static/js/editorjscode';
import ImageTool from '@editorjs/image';

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
                image: {
                    class: ImageTool,
                    config: {
                        uploader: {

                            uploadByFile(file){
                                console.log(file)
                                return {
                                    success: 1,
                                    file: {
                                        url: "/static/img/blog/"+file.name,
                                    }
                                };
                            },

                            uploadByUrl(url){
                                return {
                                    success: 1,
                                    file: {
                                        url: url,
                                    }
                                }
                            }
                        }
                    },
                    inlineToolbar: true
                },
                header:Header,
                delimiter: Delimiter,
                paragraph: {
                    class: Paragraph,
                    inlineToolbar: true,
                },
                code: CodeTool,
                embed: Embed,
                list: {
                    class: List,
                    inlineToolbar: true,
                },
            },
            data:data,
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